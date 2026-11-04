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

package v20260520

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2026-05-20"

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


func NewCallMCPToolRequest() (request *CallMCPToolRequest) {
    request = &CallMCPToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CallMCPTool")
    
    
    return
}

func NewCallMCPToolResponse() (response *CallMCPToolResponse) {
    response = &CallMCPToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CallMCPTool
// Tool call triggered by a user in the MCP App iframe.
func (c *Client) CallMCPTool(request *CallMCPToolRequest) (response *CallMCPToolResponse, err error) {
    return c.CallMCPToolWithContext(context.Background(), request)
}

// CallMCPTool
// Tool call triggered by a user in the MCP App iframe.
func (c *Client) CallMCPToolWithContext(ctx context.Context, request *CallMCPToolRequest) (response *CallMCPToolResponse, err error) {
    if request == nil {
        request = NewCallMCPToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CallMCPTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallMCPTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallMCPToolResponse()
    err = c.Send(request, response)
    return
}

func NewCheckLabelRequest() (request *CheckLabelRequest) {
    request = &CheckLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CheckLabel")
    
    
    return
}

func NewCheckLabelResponse() (response *CheckLabelResponse) {
    response = &CheckLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckLabel
// Validate tags
func (c *Client) CheckLabel(request *CheckLabelRequest) (response *CheckLabelResponse, err error) {
    return c.CheckLabelWithContext(context.Background(), request)
}

// CheckLabel
// Validate tags
func (c *Client) CheckLabelWithContext(ctx context.Context, request *CheckLabelRequest) (response *CheckLabelResponse, err error) {
    if request == nil {
        request = NewCheckLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CheckLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAgentFromAppRequest() (request *CopyAgentFromAppRequest) {
    request = &CopyAgentFromAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CopyAgentFromApp")
    
    
    return
}

func NewCopyAgentFromAppResponse() (response *CopyAgentFromAppResponse) {
    response = &CopyAgentFromAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyAgentFromApp
// This API is used to replicate an Agent. Currently only support claw mode.
func (c *Client) CopyAgentFromApp(request *CopyAgentFromAppRequest) (response *CopyAgentFromAppResponse, err error) {
    return c.CopyAgentFromAppWithContext(context.Background(), request)
}

// CopyAgentFromApp
// This API is used to replicate an Agent. Currently only support claw mode.
func (c *Client) CopyAgentFromAppWithContext(ctx context.Context, request *CopyAgentFromAppRequest) (response *CopyAgentFromAppResponse, err error) {
    if request == nil {
        request = NewCopyAgentFromAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CopyAgentFromApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyAgentFromApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyAgentFromAppResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAppRequest() (request *CopyAppRequest) {
    request = &CopyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CopyApp")
    
    
    return
}

func NewCopyAppResponse() (response *CopyAppResponse) {
    response = &CopyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyApp
// Copy an application
func (c *Client) CopyApp(request *CopyAppRequest) (response *CopyAppResponse, err error) {
    return c.CopyAppWithContext(context.Background(), request)
}

// CopyApp
// Copy an application
func (c *Client) CopyAppWithContext(ctx context.Context, request *CopyAppRequest) (response *CopyAppResponse, err error) {
    if request == nil {
        request = NewCopyAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CopyApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentRequest() (request *CreateAgentRequest) {
    request = &CreateAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateAgent")
    
    
    return
}

func NewCreateAgentResponse() (response *CreateAgentResponse) {
    response = &CreateAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgent
// This API is used to create an Agent. Currently only support claw mode.
func (c *Client) CreateAgent(request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    return c.CreateAgentWithContext(context.Background(), request)
}

// CreateAgent
// This API is used to create an Agent. Currently only support claw mode.
func (c *Client) CreateAgentWithContext(ctx context.Context, request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    if request == nil {
        request = NewCreateAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApp
// This API is used to create an application.
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// This API is used to create an application.
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppTriggerRequest() (request *CreateAppTriggerRequest) {
    request = &CreateAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateAppTrigger")
    
    
    return
}

func NewCreateAppTriggerResponse() (response *CreateAppTriggerResponse) {
    response = &CreateAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAppTrigger
// CreateAppTrigger
func (c *Client) CreateAppTrigger(request *CreateAppTriggerRequest) (response *CreateAppTriggerResponse, err error) {
    return c.CreateAppTriggerWithContext(context.Background(), request)
}

// CreateAppTrigger
// CreateAppTrigger
func (c *Client) CreateAppTriggerWithContext(ctx context.Context, request *CreateAppTriggerRequest) (response *CreateAppTriggerResponse, err error) {
    if request == nil {
        request = NewCreateAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCategoryRequest() (request *CreateCategoryRequest) {
    request = &CreateCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateCategory")
    
    
    return
}

func NewCreateCategoryResponse() (response *CreateCategoryResponse) {
    response = &CreateCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCategory
// Create a category
func (c *Client) CreateCategory(request *CreateCategoryRequest) (response *CreateCategoryResponse, err error) {
    return c.CreateCategoryWithContext(context.Background(), request)
}

// CreateCategory
// Create a category
func (c *Client) CreateCategoryWithContext(ctx context.Context, request *CreateCategoryRequest) (response *CreateCategoryResponse, err error) {
    if request == nil {
        request = NewCreateCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelRequest() (request *CreateChannelRequest) {
    request = &CreateChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateChannel")
    
    
    return
}

func NewCreateChannelResponse() (response *CreateChannelResponse) {
    response = &CreateChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChannel
// Create a channel (distinguish B-end application release channels from C-end IM channels by scene)
func (c *Client) CreateChannel(request *CreateChannelRequest) (response *CreateChannelResponse, err error) {
    return c.CreateChannelWithContext(context.Background(), request)
}

// CreateChannel
// Create a channel (distinguish B-end application release channels from C-end IM channels by scene)
func (c *Client) CreateChannelWithContext(ctx context.Context, request *CreateChannelRequest) (response *CreateChannelResponse, err error) {
    if request == nil {
        request = NewCreateChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConversationRequest() (request *CreateConversationRequest) {
    request = &CreateConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateConversation")
    
    
    return
}

func NewCreateConversationResponse() (response *CreateConversationResponse) {
    response = &CreateConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConversation
// Create session
func (c *Client) CreateConversation(request *CreateConversationRequest) (response *CreateConversationResponse, err error) {
    return c.CreateConversationWithContext(context.Background(), request)
}

// CreateConversation
// Create session
func (c *Client) CreateConversationWithContext(ctx context.Context, request *CreateConversationRequest) (response *CreateConversationResponse, err error) {
    if request == nil {
        request = NewCreateConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConversationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKBRequest() (request *CreateKBRequest) {
    request = &CreateKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateKB")
    
    
    return
}

func NewCreateKBResponse() (response *CreateKBResponse) {
    response = &CreateKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKB
// This API is used to create a knowledge base.
func (c *Client) CreateKB(request *CreateKBRequest) (response *CreateKBResponse, err error) {
    return c.CreateKBWithContext(context.Background(), request)
}

// CreateKB
// This API is used to create a knowledge base.
func (c *Client) CreateKBWithContext(ctx context.Context, request *CreateKBRequest) (response *CreateKBResponse, err error) {
    if request == nil {
        request = NewCreateKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKBResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLabelRequest() (request *CreateLabelRequest) {
    request = &CreateLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateLabel")
    
    
    return
}

func NewCreateLabelResponse() (response *CreateLabelResponse) {
    response = &CreateLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLabel
// Create a tag
func (c *Client) CreateLabel(request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
    return c.CreateLabelWithContext(context.Background(), request)
}

// CreateLabel
// Create a tag
func (c *Client) CreateLabelWithContext(ctx context.Context, request *CreateLabelRequest) (response *CreateLabelResponse, err error) {
    if request == nil {
        request = NewCreateLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMsgRecordCategoryRequest() (request *CreateMsgRecordCategoryRequest) {
    request = &CreateMsgRecordCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateMsgRecordCategory")
    
    
    return
}

func NewCreateMsgRecordCategoryResponse() (response *CreateMsgRecordCategoryResponse) {
    response = &CreateMsgRecordCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMsgRecordCategory
// This API is used to create a message record category, with support for specifying the category name and parent category. A ParentId of 0 indicates a level-1 category.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMsgRecordCategory(request *CreateMsgRecordCategoryRequest) (response *CreateMsgRecordCategoryResponse, err error) {
    return c.CreateMsgRecordCategoryWithContext(context.Background(), request)
}

// CreateMsgRecordCategory
// This API is used to create a message record category, with support for specifying the category name and parent category. A ParentId of 0 indicates a level-1 category.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMsgRecordCategoryWithContext(ctx context.Context, request *CreateMsgRecordCategoryRequest) (response *CreateMsgRecordCategoryResponse, err error) {
    if request == nil {
        request = NewCreateMsgRecordCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateMsgRecordCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMsgRecordCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMsgRecordCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePluginRequest() (request *CreatePluginRequest) {
    request = &CreatePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreatePlugin")
    
    
    return
}

func NewCreatePluginResponse() (response *CreatePluginResponse) {
    response = &CreatePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlugin
// This API is used to query plugin details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePlugin(request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    return c.CreatePluginWithContext(context.Background(), request)
}

// CreatePlugin
// This API is used to query plugin details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePluginWithContext(ctx context.Context, request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    if request == nil {
        request = NewCreatePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreatePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQAGenerationTaskRequest() (request *CreateQAGenerationTaskRequest) {
    request = &CreateQAGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateQAGenerationTask")
    
    
    return
}

func NewCreateQAGenerationTaskResponse() (response *CreateQAGenerationTaskResponse) {
    response = &CreateQAGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQAGenerationTask
// Create a QA generation task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAGenerationTask(request *CreateQAGenerationTaskRequest) (response *CreateQAGenerationTaskResponse, err error) {
    return c.CreateQAGenerationTaskWithContext(context.Background(), request)
}

// CreateQAGenerationTask
// Create a QA generation task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAGenerationTaskWithContext(ctx context.Context, request *CreateQAGenerationTaskRequest) (response *CreateQAGenerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateQAGenerationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateQAGenerationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQAGenerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQAListRequest() (request *CreateQAListRequest) {
    request = &CreateQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateQAList")
    
    
    return
}

func NewCreateQAListResponse() (response *CreateQAListResponse) {
    response = &CreateQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQAList
// Batch create QA
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAList(request *CreateQAListRequest) (response *CreateQAListResponse, err error) {
    return c.CreateQAListWithContext(context.Background(), request)
}

// CreateQAList
// Batch create QA
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateQAListWithContext(ctx context.Context, request *CreateQAListRequest) (response *CreateQAListResponse, err error) {
    if request == nil {
        request = NewCreateQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseRequest() (request *CreateReleaseRequest) {
    request = &CreateReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateRelease")
    
    
    return
}

func NewCreateReleaseResponse() (response *CreateReleaseResponse) {
    response = &CreateReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRelease
// Add new release task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    return c.CreateReleaseWithContext(context.Background(), request)
}

// CreateRelease
// Add new release task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateReleaseWithContext(ctx context.Context, request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    if request == nil {
        request = NewCreateReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSimilarQuestionRequest() (request *CreateSimilarQuestionRequest) {
    request = &CreateSimilarQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSimilarQuestion")
    
    
    return
}

func NewCreateSimilarQuestionResponse() (response *CreateSimilarQuestionResponse) {
    response = &CreateSimilarQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSimilarQuestion
// This API is used to create a similar question generation task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSimilarQuestion(request *CreateSimilarQuestionRequest) (response *CreateSimilarQuestionResponse, err error) {
    return c.CreateSimilarQuestionWithContext(context.Background(), request)
}

// CreateSimilarQuestion
// This API is used to create a similar question generation task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSimilarQuestionWithContext(ctx context.Context, request *CreateSimilarQuestionRequest) (response *CreateSimilarQuestionResponse, err error) {
    if request == nil {
        request = NewCreateSimilarQuestionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSimilarQuestion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSimilarQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSimilarQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSkillRequest() (request *CreateSkillRequest) {
    request = &CreateSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSkill")
    
    
    return
}

func NewCreateSkillResponse() (response *CreateSkillResponse) {
    response = &CreateSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSkill
// Creates a skill
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSkill(request *CreateSkillRequest) (response *CreateSkillResponse, err error) {
    return c.CreateSkillWithContext(context.Background(), request)
}

// CreateSkill
// Creates a skill
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSkillWithContext(ctx context.Context, request *CreateSkillRequest) (response *CreateSkillResponse, err error) {
    if request == nil {
        request = NewCreateSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSkillResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSkillShareRequest() (request *CreateSkillShareRequest) {
    request = &CreateSkillShareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSkillShare")
    
    
    return
}

func NewCreateSkillShareResponse() (response *CreateSkillShareResponse) {
    response = &CreateSkillShareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSkillShare
// Submit a custom Skill for enterprise-level shared approval (two-segment: submit, approve, callback to create a shared task)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSkillShare(request *CreateSkillShareRequest) (response *CreateSkillShareResponse, err error) {
    return c.CreateSkillShareWithContext(context.Background(), request)
}

// CreateSkillShare
// Submit a custom Skill for enterprise-level shared approval (two-segment: submit, approve, callback to create a shared task)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSkillShareWithContext(ctx context.Context, request *CreateSkillShareRequest) (response *CreateSkillShareResponse, err error) {
    if request == nil {
        request = NewCreateSkillShareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSkillShare")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSkillShare require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSkillShareResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSpaceRequest() (request *CreateSpaceRequest) {
    request = &CreateSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateSpace")
    
    
    return
}

func NewCreateSpaceResponse() (response *CreateSpaceResponse) {
    response = &CreateSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSpace
// Create a space
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSpace(request *CreateSpaceRequest) (response *CreateSpaceResponse, err error) {
    return c.CreateSpaceWithContext(context.Background(), request)
}

// CreateSpace
// Create a space
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSpaceWithContext(ctx context.Context, request *CreateSpaceRequest) (response *CreateSpaceResponse, err error) {
    if request == nil {
        request = NewCreateSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVariableRequest() (request *CreateVariableRequest) {
    request = &CreateVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateVariable")
    
    
    return
}

func NewCreateVariableResponse() (response *CreateVariableResponse) {
    response = &CreateVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVariable
// This API is used to create parameter variables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVariable(request *CreateVariableRequest) (response *CreateVariableResponse, err error) {
    return c.CreateVariableWithContext(context.Background(), request)
}

// CreateVariable
// This API is used to create parameter variables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateVariableWithContext(ctx context.Context, request *CreateVariableRequest) (response *CreateVariableResponse, err error) {
    if request == nil {
        request = NewCreateVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVariableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebSocketTokenRequest() (request *CreateWebSocketTokenRequest) {
    request = &CreateWebSocketTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateWebSocketToken")
    
    
    return
}

func NewCreateWebSocketTokenResponse() (response *CreateWebSocketTokenResponse) {
    response = &CreateWebSocketTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebSocketToken
// This API is used to create a WebSocket Token.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWebSocketToken(request *CreateWebSocketTokenRequest) (response *CreateWebSocketTokenResponse, err error) {
    return c.CreateWebSocketTokenWithContext(context.Background(), request)
}

// CreateWebSocketToken
// This API is used to create a WebSocket Token.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWebSocketTokenWithContext(ctx context.Context, request *CreateWebSocketTokenRequest) (response *CreateWebSocketTokenResponse, err error) {
    if request == nil {
        request = NewCreateWebSocketTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateWebSocketToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebSocketToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebSocketTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRunRequest() (request *CreateWorkflowRunRequest) {
    request = &CreateWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateWorkflowRun")
    
    
    return
}

func NewCreateWorkflowRunResponse() (response *CreateWorkflowRunResponse) {
    response = &CreateWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflowRun
// This API is used to create an asynchronous running instance of a workflow. After successful creation, the workflow runs asynchronously in the backend. The API returns the workflow running instance ID (WorkflowRunId) and other information. You can call the DescribeWorkflowRun API to query the workflow running details.
//
// Note: Asynchronous execution of a workflow is application-based. The corresponding application configuration must be set to "Single Workflow Mode" before it can be successfully created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowRun(request *CreateWorkflowRunRequest) (response *CreateWorkflowRunResponse, err error) {
    return c.CreateWorkflowRunWithContext(context.Background(), request)
}

// CreateWorkflowRun
// This API is used to create an asynchronous running instance of a workflow. After successful creation, the workflow runs asynchronously in the backend. The API returns the workflow running instance ID (WorkflowRunId) and other information. You can call the DescribeWorkflowRun API to query the workflow running details.
//
// Note: Asynchronous execution of a workflow is application-based. The corresponding application configuration must be set to "Single Workflow Mode" before it can be successfully created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowRunWithContext(ctx context.Context, request *CreateWorkflowRunRequest) (response *CreateWorkflowRunResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceCredentialRequest() (request *CreateWorkspaceCredentialRequest) {
    request = &CreateWorkspaceCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "CreateWorkspaceCredential")
    
    
    return
}

func NewCreateWorkspaceCredentialResponse() (response *CreateWorkspaceCredentialResponse) {
    response = &CreateWorkspaceCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspaceCredential
// Create a workspace credential
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkspaceCredential(request *CreateWorkspaceCredentialRequest) (response *CreateWorkspaceCredentialResponse, err error) {
    return c.CreateWorkspaceCredentialWithContext(context.Background(), request)
}

// CreateWorkspaceCredential
// Create a workspace credential
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkspaceCredentialWithContext(ctx context.Context, request *CreateWorkspaceCredentialRequest) (response *CreateWorkspaceCredentialResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "CreateWorkspaceCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentRequest() (request *DeleteAgentRequest) {
    request = &DeleteAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteAgent")
    
    
    return
}

func NewDeleteAgentResponse() (response *DeleteAgentResponse) {
    response = &DeleteAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgent
// This API is used to delete an Agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAgent(request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
    return c.DeleteAgentWithContext(context.Background(), request)
}

// DeleteAgent
// This API is used to delete an Agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAgentWithContext(ctx context.Context, request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
    if request == nil {
        request = NewDeleteAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
    request = &DeleteAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteApp")
    
    
    return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
    response = &DeleteAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApp
// This API is used to delete applications
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    return c.DeleteAppWithContext(context.Background(), request)
}

// DeleteApp
// This API is used to delete applications
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    if request == nil {
        request = NewDeleteAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppTriggerRequest() (request *DeleteAppTriggerRequest) {
    request = &DeleteAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteAppTrigger")
    
    
    return
}

func NewDeleteAppTriggerResponse() (response *DeleteAppTriggerResponse) {
    response = &DeleteAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAppTrigger
// DeleteAppTrigger
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAppTrigger(request *DeleteAppTriggerRequest) (response *DeleteAppTriggerResponse, err error) {
    return c.DeleteAppTriggerWithContext(context.Background(), request)
}

// DeleteAppTrigger
// DeleteAppTrigger
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAppTriggerWithContext(ctx context.Context, request *DeleteAppTriggerRequest) (response *DeleteAppTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCategoryRequest() (request *DeleteCategoryRequest) {
    request = &DeleteCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteCategory")
    
    
    return
}

func NewDeleteCategoryResponse() (response *DeleteCategoryResponse) {
    response = &DeleteCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCategory
// Delete a category
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
    return c.DeleteCategoryWithContext(context.Background(), request)
}

// DeleteCategory
// Delete a category
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteCategoryWithContext(ctx context.Context, request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
    if request == nil {
        request = NewDeleteCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteChannelRequest() (request *DeleteChannelRequest) {
    request = &DeleteChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteChannel")
    
    
    return
}

func NewDeleteChannelResponse() (response *DeleteChannelResponse) {
    response = &DeleteChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteChannel
// Delete a channel (scenarios are distinguished by scene)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteChannel(request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    return c.DeleteChannelWithContext(context.Background(), request)
}

// DeleteChannel
// Delete a channel (scenarios are distinguished by scene)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteChannelWithContext(ctx context.Context, request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    if request == nil {
        request = NewDeleteChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConversationRequest() (request *DeleteConversationRequest) {
    request = &DeleteConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteConversation")
    
    
    return
}

func NewDeleteConversationResponse() (response *DeleteConversationResponse) {
    response = &DeleteConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConversation
// Delete conversation
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteConversation(request *DeleteConversationRequest) (response *DeleteConversationResponse, err error) {
    return c.DeleteConversationWithContext(context.Background(), request)
}

// DeleteConversation
// Delete conversation
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteConversationWithContext(ctx context.Context, request *DeleteConversationRequest) (response *DeleteConversationResponse, err error) {
    if request == nil {
        request = NewDeleteConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocListRequest() (request *DeleteDocListRequest) {
    request = &DeleteDocListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteDocList")
    
    
    return
}

func NewDeleteDocListResponse() (response *DeleteDocListResponse) {
    response = &DeleteDocListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDocList
// Delete documents in batches
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteDocList(request *DeleteDocListRequest) (response *DeleteDocListResponse, err error) {
    return c.DeleteDocListWithContext(context.Background(), request)
}

// DeleteDocList
// Delete documents in batches
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteDocListWithContext(ctx context.Context, request *DeleteDocListRequest) (response *DeleteDocListResponse, err error) {
    if request == nil {
        request = NewDeleteDocListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteDocList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKBRequest() (request *DeleteKBRequest) {
    request = &DeleteKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteKB")
    
    
    return
}

func NewDeleteKBResponse() (response *DeleteKBResponse) {
    response = &DeleteKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKB
// This API is used to delete the knowledge base.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteKB(request *DeleteKBRequest) (response *DeleteKBResponse, err error) {
    return c.DeleteKBWithContext(context.Background(), request)
}

// DeleteKB
// This API is used to delete the knowledge base.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteKBWithContext(ctx context.Context, request *DeleteKBRequest) (response *DeleteKBResponse, err error) {
    if request == nil {
        request = NewDeleteKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKBResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLabelListRequest() (request *DeleteLabelListRequest) {
    request = &DeleteLabelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteLabelList")
    
    
    return
}

func NewDeleteLabelListResponse() (response *DeleteLabelListResponse) {
    response = &DeleteLabelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLabelList
// Batch deletes tags
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLabelList(request *DeleteLabelListRequest) (response *DeleteLabelListResponse, err error) {
    return c.DeleteLabelListWithContext(context.Background(), request)
}

// DeleteLabelList
// Batch deletes tags
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLabelListWithContext(ctx context.Context, request *DeleteLabelListRequest) (response *DeleteLabelListResponse, err error) {
    if request == nil {
        request = NewDeleteLabelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteLabelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLabelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLabelListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMsgRecordCategoryRequest() (request *DeleteMsgRecordCategoryRequest) {
    request = &DeleteMsgRecordCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteMsgRecordCategory")
    
    
    return
}

func NewDeleteMsgRecordCategoryResponse() (response *DeleteMsgRecordCategoryResponse) {
    response = &DeleteMsgRecordCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMsgRecordCategory
// Deletes a specified message record category
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMsgRecordCategory(request *DeleteMsgRecordCategoryRequest) (response *DeleteMsgRecordCategoryResponse, err error) {
    return c.DeleteMsgRecordCategoryWithContext(context.Background(), request)
}

// DeleteMsgRecordCategory
// Deletes a specified message record category
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMsgRecordCategoryWithContext(ctx context.Context, request *DeleteMsgRecordCategoryRequest) (response *DeleteMsgRecordCategoryResponse, err error) {
    if request == nil {
        request = NewDeleteMsgRecordCategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteMsgRecordCategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMsgRecordCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMsgRecordCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePluginRequest() (request *DeletePluginRequest) {
    request = &DeletePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeletePlugin")
    
    
    return
}

func NewDeletePluginResponse() (response *DeletePluginResponse) {
    response = &DeletePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePlugin
// Modify a plug-in
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    return c.DeletePluginWithContext(context.Background(), request)
}

// DeletePlugin
// Modify a plug-in
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePluginWithContext(ctx context.Context, request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    if request == nil {
        request = NewDeletePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeletePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQAListRequest() (request *DeleteQAListRequest) {
    request = &DeleteQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteQAList")
    
    
    return
}

func NewDeleteQAListResponse() (response *DeleteQAListResponse) {
    response = &DeleteQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQAList
// Delete QAs in batches
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteQAList(request *DeleteQAListRequest) (response *DeleteQAListResponse, err error) {
    return c.DeleteQAListWithContext(context.Background(), request)
}

// DeleteQAList
// Delete QAs in batches
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteQAListWithContext(ctx context.Context, request *DeleteQAListRequest) (response *DeleteQAListResponse, err error) {
    if request == nil {
        request = NewDeleteQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSkillRequest() (request *DeleteSkillRequest) {
    request = &DeleteSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteSkill")
    
    
    return
}

func NewDeleteSkillResponse() (response *DeleteSkillResponse) {
    response = &DeleteSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSkill
// Delete a Custom Skill Authorization: creator ∨ (edit permission ∧ delete permission) Denied scenarios: non-Custom type / shared / in security detection / pending listing approval / pending delisting approval
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSkill(request *DeleteSkillRequest) (response *DeleteSkillResponse, err error) {
    return c.DeleteSkillWithContext(context.Background(), request)
}

// DeleteSkill
// Delete a Custom Skill Authorization: creator ∨ (edit permission ∧ delete permission) Denied scenarios: non-Custom type / shared / in security detection / pending listing approval / pending delisting approval
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSkillWithContext(ctx context.Context, request *DeleteSkillRequest) (response *DeleteSkillResponse, err error) {
    if request == nil {
        request = NewDeleteSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSkillResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSkillShareRequest() (request *DeleteSkillShareRequest) {
    request = &DeleteSkillShareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteSkillShare")
    
    
    return
}

func NewDeleteSkillShareResponse() (response *DeleteSkillShareResponse) {
    response = &DeleteSkillShareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSkillShare
// Submit a shared Skill delisting approval (v2, two-segment: submit → approve → callback to delist the shared Skill) Authorization: right to erasure Denied scenarios: not shared / pending listing approval / pending delisting approval
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSkillShare(request *DeleteSkillShareRequest) (response *DeleteSkillShareResponse, err error) {
    return c.DeleteSkillShareWithContext(context.Background(), request)
}

// DeleteSkillShare
// Submit a shared Skill delisting approval (v2, two-segment: submit → approve → callback to delist the shared Skill) Authorization: right to erasure Denied scenarios: not shared / pending listing approval / pending delisting approval
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSkillShareWithContext(ctx context.Context, request *DeleteSkillShareRequest) (response *DeleteSkillShareResponse, err error) {
    if request == nil {
        request = NewDeleteSkillShareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteSkillShare")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSkillShare require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSkillShareResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSpaceRequest() (request *DeleteSpaceRequest) {
    request = &DeleteSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteSpace")
    
    
    return
}

func NewDeleteSpaceResponse() (response *DeleteSpaceResponse) {
    response = &DeleteSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSpace
// Delete a space
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSpace(request *DeleteSpaceRequest) (response *DeleteSpaceResponse, err error) {
    return c.DeleteSpaceWithContext(context.Background(), request)
}

// DeleteSpace
// Delete a space
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSpaceWithContext(ctx context.Context, request *DeleteSpaceRequest) (response *DeleteSpaceResponse, err error) {
    if request == nil {
        request = NewDeleteSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVariableRequest() (request *DeleteVariableRequest) {
    request = &DeleteVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DeleteVariable")
    
    
    return
}

func NewDeleteVariableResponse() (response *DeleteVariableResponse) {
    response = &DeleteVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVariable
// Delete a parameter variable
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVariable(request *DeleteVariableRequest) (response *DeleteVariableResponse, err error) {
    return c.DeleteVariableWithContext(context.Background(), request)
}

// DeleteVariable
// Delete a parameter variable
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVariableWithContext(ctx context.Context, request *DeleteVariableRequest) (response *DeleteVariableResponse, err error) {
    if request == nil {
        request = NewDeleteVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DeleteVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVariableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountListRequest() (request *DescribeAccountListRequest) {
    request = &DescribeAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAccountList")
    
    
    return
}

func NewDescribeAccountListResponse() (response *DescribeAccountListResponse) {
    response = &DescribeAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountList
// View the employee list of a business
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAccountList(request *DescribeAccountListRequest) (response *DescribeAccountListResponse, err error) {
    return c.DescribeAccountListWithContext(context.Background(), request)
}

// DescribeAccountList
// View the employee list of a business
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAccountListWithContext(ctx context.Context, request *DescribeAccountListRequest) (response *DescribeAccountListResponse, err error) {
    if request == nil {
        request = NewDescribeAccountListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAccountList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentDetailRequest() (request *DescribeAgentDetailRequest) {
    request = &DescribeAgentDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAgentDetail")
    
    
    return
}

func NewDescribeAgentDetailResponse() (response *DescribeAgentDetailResponse) {
    response = &DescribeAgentDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentDetail
// Query Agent details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAgentDetail(request *DescribeAgentDetailRequest) (response *DescribeAgentDetailResponse, err error) {
    return c.DescribeAgentDetailWithContext(context.Background(), request)
}

// DescribeAgentDetail
// Query Agent details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAgentDetailWithContext(ctx context.Context, request *DescribeAgentDetailRequest) (response *DescribeAgentDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAgentDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAgentDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentReleasePreviewListRequest() (request *DescribeAgentReleasePreviewListRequest) {
    request = &DescribeAgentReleasePreviewListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAgentReleasePreviewList")
    
    
    return
}

func NewDescribeAgentReleasePreviewListResponse() (response *DescribeAgentReleasePreviewListResponse) {
    response = &DescribeAgentReleasePreviewListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentReleasePreviewList
// Get the release preview list of agents under an application
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentReleasePreviewList(request *DescribeAgentReleasePreviewListRequest) (response *DescribeAgentReleasePreviewListResponse, err error) {
    return c.DescribeAgentReleasePreviewListWithContext(context.Background(), request)
}

// DescribeAgentReleasePreviewList
// Get the release preview list of agents under an application
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentReleasePreviewListWithContext(ctx context.Context, request *DescribeAgentReleasePreviewListRequest) (response *DescribeAgentReleasePreviewListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentReleasePreviewListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAgentReleasePreviewList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentReleasePreviewList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentReleasePreviewListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentSummaryListRequest() (request *DescribeAgentSummaryListRequest) {
    request = &DescribeAgentSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAgentSummaryList")
    
    
    return
}

func NewDescribeAgentSummaryListResponse() (response *DescribeAgentSummaryListResponse) {
    response = &DescribeAgentSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentSummaryList
// Queries the abstract information list of an Agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentSummaryList(request *DescribeAgentSummaryListRequest) (response *DescribeAgentSummaryListResponse, err error) {
    return c.DescribeAgentSummaryListWithContext(context.Background(), request)
}

// DescribeAgentSummaryList
// Queries the abstract information list of an Agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAgentSummaryListWithContext(ctx context.Context, request *DescribeAgentSummaryListRequest) (response *DescribeAgentSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAgentSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppRequest() (request *DescribeAppRequest) {
    request = &DescribeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeApp")
    
    
    return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
    response = &DescribeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApp
// This API is used to obtain application information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    return c.DescribeAppWithContext(context.Background(), request)
}

// DescribeApp
// This API is used to obtain application information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppWithContext(ctx context.Context, request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    if request == nil {
        request = NewDescribeAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppStatisticsOverviewRequest() (request *DescribeAppStatisticsOverviewRequest) {
    request = &DescribeAppStatisticsOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppStatisticsOverview")
    
    
    return
}

func NewDescribeAppStatisticsOverviewResponse() (response *DescribeAppStatisticsOverviewResponse) {
    response = &DescribeAppStatisticsOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppStatisticsOverview
// This API is used to query the invocation statistics overview under an application view, including the total number of calls, call success rate, avg duration of total tokens, and avg duration of first tokens. For RAG applications, it also returns the count and proportion of calls for each response method for drawing a pie chart.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppStatisticsOverview(request *DescribeAppStatisticsOverviewRequest) (response *DescribeAppStatisticsOverviewResponse, err error) {
    return c.DescribeAppStatisticsOverviewWithContext(context.Background(), request)
}

// DescribeAppStatisticsOverview
// This API is used to query the invocation statistics overview under an application view, including the total number of calls, call success rate, avg duration of total tokens, and avg duration of first tokens. For RAG applications, it also returns the count and proportion of calls for each response method for drawing a pie chart.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppStatisticsOverviewWithContext(ctx context.Context, request *DescribeAppStatisticsOverviewRequest) (response *DescribeAppStatisticsOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppStatisticsOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppStatisticsOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppStatisticsOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppSummaryListRequest() (request *DescribeAppSummaryListRequest) {
    request = &DescribeAppSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppSummaryList")
    
    
    return
}

func NewDescribeAppSummaryListResponse() (response *DescribeAppSummaryListResponse) {
    response = &DescribeAppSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppSummaryList
// Retrieve the abstract list of applications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppSummaryList(request *DescribeAppSummaryListRequest) (response *DescribeAppSummaryListResponse, err error) {
    return c.DescribeAppSummaryListWithContext(context.Background(), request)
}

// DescribeAppSummaryList
// Retrieve the abstract list of applications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppSummaryListWithContext(ctx context.Context, request *DescribeAppSummaryListRequest) (response *DescribeAppSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeAppSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerRequest() (request *DescribeAppTriggerRequest) {
    request = &DescribeAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTrigger")
    
    
    return
}

func NewDescribeAppTriggerResponse() (response *DescribeAppTriggerResponse) {
    response = &DescribeAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTrigger
// DescribeAppTrigger
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTrigger(request *DescribeAppTriggerRequest) (response *DescribeAppTriggerResponse, err error) {
    return c.DescribeAppTriggerWithContext(context.Background(), request)
}

// DescribeAppTrigger
// DescribeAppTrigger
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerWithContext(ctx context.Context, request *DescribeAppTriggerRequest) (response *DescribeAppTriggerResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerInstanceRequest() (request *DescribeAppTriggerInstanceRequest) {
    request = &DescribeAppTriggerInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTriggerInstance")
    
    
    return
}

func NewDescribeAppTriggerInstanceResponse() (response *DescribeAppTriggerInstanceResponse) {
    response = &DescribeAppTriggerInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTriggerInstance
// DescribeAppTriggerInstance
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerInstance(request *DescribeAppTriggerInstanceRequest) (response *DescribeAppTriggerInstanceResponse, err error) {
    return c.DescribeAppTriggerInstanceWithContext(context.Background(), request)
}

// DescribeAppTriggerInstance
// DescribeAppTriggerInstance
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerInstanceWithContext(ctx context.Context, request *DescribeAppTriggerInstanceRequest) (response *DescribeAppTriggerInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTriggerInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTriggerInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerRunLogListRequest() (request *DescribeAppTriggerRunLogListRequest) {
    request = &DescribeAppTriggerRunLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTriggerRunLogList")
    
    
    return
}

func NewDescribeAppTriggerRunLogListResponse() (response *DescribeAppTriggerRunLogListResponse) {
    response = &DescribeAppTriggerRunLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTriggerRunLogList
// DescribeAppTriggerRunLogList
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerRunLogList(request *DescribeAppTriggerRunLogListRequest) (response *DescribeAppTriggerRunLogListResponse, err error) {
    return c.DescribeAppTriggerRunLogListWithContext(context.Background(), request)
}

// DescribeAppTriggerRunLogList
// DescribeAppTriggerRunLogList
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerRunLogListWithContext(ctx context.Context, request *DescribeAppTriggerRunLogListRequest) (response *DescribeAppTriggerRunLogListResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerRunLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTriggerRunLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTriggerRunLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerRunLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppTriggerSummaryListRequest() (request *DescribeAppTriggerSummaryListRequest) {
    request = &DescribeAppTriggerSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAppTriggerSummaryList")
    
    
    return
}

func NewDescribeAppTriggerSummaryListResponse() (response *DescribeAppTriggerSummaryListResponse) {
    response = &DescribeAppTriggerSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppTriggerSummaryList
// DescribeAppTriggerSummaryList
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerSummaryList(request *DescribeAppTriggerSummaryListRequest) (response *DescribeAppTriggerSummaryListResponse, err error) {
    return c.DescribeAppTriggerSummaryListWithContext(context.Background(), request)
}

// DescribeAppTriggerSummaryList
// DescribeAppTriggerSummaryList
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAppTriggerSummaryListWithContext(ctx context.Context, request *DescribeAppTriggerSummaryListRequest) (response *DescribeAppTriggerSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeAppTriggerSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAppTriggerSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppTriggerSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppTriggerSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogListRequest() (request *DescribeAuditLogListRequest) {
    request = &DescribeAuditLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeAuditLogList")
    
    
    return
}

func NewDescribeAuditLogListResponse() (response *DescribeAuditLogListResponse) {
    response = &DescribeAuditLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogList
// View operation log list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditLogList(request *DescribeAuditLogListRequest) (response *DescribeAuditLogListResponse, err error) {
    return c.DescribeAuditLogListWithContext(context.Background(), request)
}

// DescribeAuditLogList
// View operation log list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditLogListWithContext(ctx context.Context, request *DescribeAuditLogListRequest) (response *DescribeAuditLogListResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeAuditLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCategoryListRequest() (request *DescribeCategoryListRequest) {
    request = &DescribeCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeCategoryList")
    
    
    return
}

func NewDescribeCategoryListResponse() (response *DescribeCategoryListResponse) {
    response = &DescribeCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCategoryList
// Query the classification list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCategoryList(request *DescribeCategoryListRequest) (response *DescribeCategoryListResponse, err error) {
    return c.DescribeCategoryListWithContext(context.Background(), request)
}

// DescribeCategoryList
// Query the classification list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCategoryListWithContext(ctx context.Context, request *DescribeCategoryListRequest) (response *DescribeCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeCategoryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeCategoryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyLimitDetailListRequest() (request *DescribeConcurrencyLimitDetailListRequest) {
    request = &DescribeConcurrencyLimitDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConcurrencyLimitDetailList")
    
    
    return
}

func NewDescribeConcurrencyLimitDetailListResponse() (response *DescribeConcurrencyLimitDetailListResponse) {
    response = &DescribeConcurrencyLimitDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyLimitDetailList
// Query concurrency overrun details, including QPM/TPM overruns and exclusive concurrency overruns. Return the overrun time, space, app, model, and request content.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConcurrencyLimitDetailList(request *DescribeConcurrencyLimitDetailListRequest) (response *DescribeConcurrencyLimitDetailListResponse, err error) {
    return c.DescribeConcurrencyLimitDetailListWithContext(context.Background(), request)
}

// DescribeConcurrencyLimitDetailList
// Query concurrency overrun details, including QPM/TPM overruns and exclusive concurrency overruns. Return the overrun time, space, app, model, and request content.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConcurrencyLimitDetailListWithContext(ctx context.Context, request *DescribeConcurrencyLimitDetailListRequest) (response *DescribeConcurrencyLimitDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyLimitDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConcurrencyLimitDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyLimitDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyLimitDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConflictQARequest() (request *DescribeConflictQARequest) {
    request = &DescribeConflictQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConflictQA")
    
    
    return
}

func NewDescribeConflictQAResponse() (response *DescribeConflictQAResponse) {
    response = &DescribeConflictQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConflictQA
// Queries the conflict query details
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQA(request *DescribeConflictQARequest) (response *DescribeConflictQAResponse, err error) {
    return c.DescribeConflictQAWithContext(context.Background(), request)
}

// DescribeConflictQA
// Queries the conflict query details
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQAWithContext(ctx context.Context, request *DescribeConflictQARequest) (response *DescribeConflictQAResponse, err error) {
    if request == nil {
        request = NewDescribeConflictQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConflictQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConflictQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConflictQAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConflictQASummaryListRequest() (request *DescribeConflictQASummaryListRequest) {
    request = &DescribeConflictQASummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConflictQASummaryList")
    
    
    return
}

func NewDescribeConflictQASummaryListResponse() (response *DescribeConflictQASummaryListResponse) {
    response = &DescribeConflictQASummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConflictQASummaryList
// Query the conflict list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQASummaryList(request *DescribeConflictQASummaryListRequest) (response *DescribeConflictQASummaryListResponse, err error) {
    return c.DescribeConflictQASummaryListWithContext(context.Background(), request)
}

// DescribeConflictQASummaryList
// Query the conflict list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConflictQASummaryListWithContext(ctx context.Context, request *DescribeConflictQASummaryListRequest) (response *DescribeConflictQASummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeConflictQASummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConflictQASummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConflictQASummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConflictQASummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationRequest() (request *DescribeConversationRequest) {
    request = &DescribeConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConversation")
    
    
    return
}

func NewDescribeConversationResponse() (response *DescribeConversationResponse) {
    response = &DescribeConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConversation
// View session information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversation(request *DescribeConversationRequest) (response *DescribeConversationResponse, err error) {
    return c.DescribeConversationWithContext(context.Background(), request)
}

// DescribeConversation
// View session information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationWithContext(ctx context.Context, request *DescribeConversationRequest) (response *DescribeConversationResponse, err error) {
    if request == nil {
        request = NewDescribeConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationListRequest() (request *DescribeConversationListRequest) {
    request = &DescribeConversationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConversationList")
    
    
    return
}

func NewDescribeConversationListResponse() (response *DescribeConversationListResponse) {
    response = &DescribeConversationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConversationList
// Get session list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationList(request *DescribeConversationListRequest) (response *DescribeConversationListResponse, err error) {
    return c.DescribeConversationListWithContext(context.Background(), request)
}

// DescribeConversationList
// Get session list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationListWithContext(ctx context.Context, request *DescribeConversationListRequest) (response *DescribeConversationListResponse, err error) {
    if request == nil {
        request = NewDescribeConversationListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConversationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConversationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConversationMessageListRequest() (request *DescribeConversationMessageListRequest) {
    request = &DescribeConversationMessageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeConversationMessageList")
    
    
    return
}

func NewDescribeConversationMessageListResponse() (response *DescribeConversationMessageListResponse) {
    response = &DescribeConversationMessageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConversationMessageList
// Get historical messages of a conversation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationMessageList(request *DescribeConversationMessageListRequest) (response *DescribeConversationMessageListResponse, err error) {
    return c.DescribeConversationMessageListWithContext(context.Background(), request)
}

// DescribeConversationMessageList
// Get historical messages of a conversation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConversationMessageListWithContext(ctx context.Context, request *DescribeConversationMessageListRequest) (response *DescribeConversationMessageListResponse, err error) {
    if request == nil {
        request = NewDescribeConversationMessageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeConversationMessageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConversationMessageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConversationMessageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLabelRequest() (request *DescribeLabelRequest) {
    request = &DescribeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeLabel")
    
    
    return
}

func NewDescribeLabelResponse() (response *DescribeLabelResponse) {
    response = &DescribeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLabel
// Queries tag details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabel(request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
    return c.DescribeLabelWithContext(context.Background(), request)
}

// DescribeLabel
// Queries tag details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabelWithContext(ctx context.Context, request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
    if request == nil {
        request = NewDescribeLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLabelSummaryListRequest() (request *DescribeLabelSummaryListRequest) {
    request = &DescribeLabelSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeLabelSummaryList")
    
    
    return
}

func NewDescribeLabelSummaryListResponse() (response *DescribeLabelSummaryListResponse) {
    response = &DescribeLabelSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLabelSummaryList
// Querying the Tag List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabelSummaryList(request *DescribeLabelSummaryListRequest) (response *DescribeLabelSummaryListResponse, err error) {
    return c.DescribeLabelSummaryListWithContext(context.Background(), request)
}

// DescribeLabelSummaryList
// Querying the Tag List
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLabelSummaryListWithContext(ctx context.Context, request *DescribeLabelSummaryListRequest) (response *DescribeLabelSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeLabelSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeLabelSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLabelSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLabelSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLatestReleaseRequest() (request *DescribeLatestReleaseRequest) {
    request = &DescribeLatestReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeLatestRelease")
    
    
    return
}

func NewDescribeLatestReleaseResponse() (response *DescribeLatestReleaseResponse) {
    response = &DescribeLatestReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLatestRelease
// Pull the latest release info (including release time, status, and channel).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestRelease(request *DescribeLatestReleaseRequest) (response *DescribeLatestReleaseResponse, err error) {
    return c.DescribeLatestReleaseWithContext(context.Background(), request)
}

// DescribeLatestRelease
// Pull the latest release info (including release time, status, and channel).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestReleaseWithContext(ctx context.Context, request *DescribeLatestReleaseRequest) (response *DescribeLatestReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeLatestReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeLatestRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLatestRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLatestReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelListRequest() (request *DescribeModelListRequest) {
    request = &DescribeModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeModelList")
    
    
    return
}

func NewDescribeModelListResponse() (response *DescribeModelListResponse) {
    response = &DescribeModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelList
// Query the model list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeModelList(request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    return c.DescribeModelListWithContext(context.Background(), request)
}

// DescribeModelList
// Query the model list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeModelListWithContext(ctx context.Context, request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    if request == nil {
        request = NewDescribeModelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeModelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsgRecordCategoryListRequest() (request *DescribeMsgRecordCategoryListRequest) {
    request = &DescribeMsgRecordCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeMsgRecordCategoryList")
    
    
    return
}

func NewDescribeMsgRecordCategoryListResponse() (response *DescribeMsgRecordCategoryListResponse) {
    response = &DescribeMsgRecordCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMsgRecordCategoryList
// Query the message record category tree of an app, return categories and subcategories, number of records under each category, and operation permissions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMsgRecordCategoryList(request *DescribeMsgRecordCategoryListRequest) (response *DescribeMsgRecordCategoryListResponse, err error) {
    return c.DescribeMsgRecordCategoryListWithContext(context.Background(), request)
}

// DescribeMsgRecordCategoryList
// Query the message record category tree of an app, return categories and subcategories, number of records under each category, and operation permissions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMsgRecordCategoryListWithContext(ctx context.Context, request *DescribeMsgRecordCategoryListRequest) (response *DescribeMsgRecordCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeMsgRecordCategoryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeMsgRecordCategoryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMsgRecordCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMsgRecordCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeBatchRunRequest() (request *DescribeNodeBatchRunRequest) {
    request = &DescribeNodeBatchRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeNodeBatchRun")
    
    
    return
}

func NewDescribeNodeBatchRunResponse() (response *DescribeNodeBatchRunResponse) {
    response = &DescribeNodeBatchRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodeBatchRun
// DescribeNodeBatchRun
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNodeBatchRun(request *DescribeNodeBatchRunRequest) (response *DescribeNodeBatchRunResponse, err error) {
    return c.DescribeNodeBatchRunWithContext(context.Background(), request)
}

// DescribeNodeBatchRun
// DescribeNodeBatchRun
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNodeBatchRunWithContext(ctx context.Context, request *DescribeNodeBatchRunRequest) (response *DescribeNodeBatchRunResponse, err error) {
    if request == nil {
        request = NewDescribeNodeBatchRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeNodeBatchRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeBatchRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeBatchRunResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeRunDetailRequest() (request *DescribeNodeRunDetailRequest) {
    request = &DescribeNodeRunDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeNodeRunDetail")
    
    
    return
}

func NewDescribeNodeRunDetailResponse() (response *DescribeNodeRunDetailResponse) {
    response = &DescribeNodeRunDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodeRunDetail
// View workflow node running details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeNodeRunDetail(request *DescribeNodeRunDetailRequest) (response *DescribeNodeRunDetailResponse, err error) {
    return c.DescribeNodeRunDetailWithContext(context.Background(), request)
}

// DescribeNodeRunDetail
// View workflow node running details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeNodeRunDetailWithContext(ctx context.Context, request *DescribeNodeRunDetailRequest) (response *DescribeNodeRunDetailResponse, err error) {
    if request == nil {
        request = NewDescribeNodeRunDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeNodeRunDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeRunDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeRunDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginRequest() (request *DescribePluginRequest) {
    request = &DescribePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribePlugin")
    
    
    return
}

func NewDescribePluginResponse() (response *DescribePluginResponse) {
    response = &DescribePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlugin
// This API is used to query plugin details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePlugin(request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    return c.DescribePluginWithContext(context.Background(), request)
}

// DescribePlugin
// This API is used to query plugin details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePluginWithContext(ctx context.Context, request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    if request == nil {
        request = NewDescribePluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribePlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginSummaryListRequest() (request *DescribePluginSummaryListRequest) {
    request = &DescribePluginSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribePluginSummaryList")
    
    
    return
}

func NewDescribePluginSummaryListResponse() (response *DescribePluginSummaryListResponse) {
    response = &DescribePluginSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePluginSummaryList
// Querying the List of Plugins
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePluginSummaryList(request *DescribePluginSummaryListRequest) (response *DescribePluginSummaryListResponse, err error) {
    return c.DescribePluginSummaryListWithContext(context.Background(), request)
}

// DescribePluginSummaryList
// Querying the List of Plugins
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePluginSummaryListWithContext(ctx context.Context, request *DescribePluginSummaryListRequest) (response *DescribePluginSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribePluginSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribePluginSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQARequest() (request *DescribeQARequest) {
    request = &DescribeQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeQA")
    
    
    return
}

func NewDescribeQAResponse() (response *DescribeQAResponse) {
    response = &DescribeQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQA
// Query QA details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQA(request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    return c.DescribeQAWithContext(context.Background(), request)
}

// DescribeQA
// Query QA details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQAWithContext(ctx context.Context, request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    if request == nil {
        request = NewDescribeQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQASummaryListRequest() (request *DescribeQASummaryListRequest) {
    request = &DescribeQASummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeQASummaryList")
    
    
    return
}

func NewDescribeQASummaryListResponse() (response *DescribeQASummaryListResponse) {
    response = &DescribeQASummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQASummaryList
// Queries the QA list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQASummaryList(request *DescribeQASummaryListRequest) (response *DescribeQASummaryListResponse, err error) {
    return c.DescribeQASummaryListWithContext(context.Background(), request)
}

// DescribeQASummaryList
// Queries the QA list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQASummaryListWithContext(ctx context.Context, request *DescribeQASummaryListRequest) (response *DescribeQASummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeQASummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeQASummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQASummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQASummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseListRequest() (request *DescribeReleaseListRequest) {
    request = &DescribeReleaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeReleaseList")
    
    
    return
}

func NewDescribeReleaseListResponse() (response *DescribeReleaseListResponse) {
    response = &DescribeReleaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseList
// Release record list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseList(request *DescribeReleaseListRequest) (response *DescribeReleaseListResponse, err error) {
    return c.DescribeReleaseListWithContext(context.Background(), request)
}

// DescribeReleaseList
// Release record list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseListWithContext(ctx context.Context, request *DescribeReleaseListRequest) (response *DescribeReleaseListResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeReleaseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseSummaryRequest() (request *DescribeReleaseSummaryRequest) {
    request = &DescribeReleaseSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeReleaseSummary")
    
    
    return
}

func NewDescribeReleaseSummaryResponse() (response *DescribeReleaseSummaryResponse) {
    response = &DescribeReleaseSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseSummary
// Queries a publication task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseSummary(request *DescribeReleaseSummaryRequest) (response *DescribeReleaseSummaryResponse, err error) {
    return c.DescribeReleaseSummaryWithContext(context.Background(), request)
}

// DescribeReleaseSummary
// Queries a publication task
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeReleaseSummaryWithContext(ctx context.Context, request *DescribeReleaseSummaryRequest) (response *DescribeReleaseSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeReleaseSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillCategoryListRequest() (request *DescribeSkillCategoryListRequest) {
    request = &DescribeSkillCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillCategoryList")
    
    
    return
}

func NewDescribeSkillCategoryListResponse() (response *DescribeSkillCategoryListResponse) {
    response = &DescribeSkillCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillCategoryList
// Query the Skill classification list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSkillCategoryList(request *DescribeSkillCategoryListRequest) (response *DescribeSkillCategoryListResponse, err error) {
    return c.DescribeSkillCategoryListWithContext(context.Background(), request)
}

// DescribeSkillCategoryList
// Query the Skill classification list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSkillCategoryListWithContext(ctx context.Context, request *DescribeSkillCategoryListRequest) (response *DescribeSkillCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillCategoryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillCategoryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillReferenceListRequest() (request *DescribeSkillReferenceListRequest) {
    request = &DescribeSkillReferenceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillReferenceList")
    
    
    return
}

func NewDescribeSkillReferenceListResponse() (response *DescribeSkillReferenceListResponse) {
    response = &DescribeSkillReferenceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillReferenceList
// This API is used to query the list of details where a certain Skill is referenced, grouped by SkillRefType: OpenClaw, cloud agent, or business assistant agent. Authentication: same as DescribeSkillDetail.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSkillReferenceList(request *DescribeSkillReferenceListRequest) (response *DescribeSkillReferenceListResponse, err error) {
    return c.DescribeSkillReferenceListWithContext(context.Background(), request)
}

// DescribeSkillReferenceList
// This API is used to query the list of details where a certain Skill is referenced, grouped by SkillRefType: OpenClaw, cloud agent, or business assistant agent. Authentication: same as DescribeSkillDetail.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSkillReferenceListWithContext(ctx context.Context, request *DescribeSkillReferenceListRequest) (response *DescribeSkillReferenceListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillReferenceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillReferenceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillReferenceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillReferenceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillSummaryListRequest() (request *DescribeSkillSummaryListRequest) {
    request = &DescribeSkillSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSkillSummaryList")
    
    
    return
}

func NewDescribeSkillSummaryListResponse() (response *DescribeSkillSummaryListResponse) {
    response = &DescribeSkillSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillSummaryList
// Query the Skill list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSkillSummaryList(request *DescribeSkillSummaryListRequest) (response *DescribeSkillSummaryListResponse, err error) {
    return c.DescribeSkillSummaryListWithContext(context.Background(), request)
}

// DescribeSkillSummaryList
// Query the Skill list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSkillSummaryListWithContext(ctx context.Context, request *DescribeSkillSummaryListRequest) (response *DescribeSkillSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSkillSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceListRequest() (request *DescribeSpaceListRequest) {
    request = &DescribeSpaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSpaceList")
    
    
    return
}

func NewDescribeSpaceListResponse() (response *DescribeSpaceListResponse) {
    response = &DescribeSpaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceList
// This API is used to obtain the space list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSpaceList(request *DescribeSpaceListRequest) (response *DescribeSpaceListResponse, err error) {
    return c.DescribeSpaceListWithContext(context.Background(), request)
}

// DescribeSpaceList
// This API is used to obtain the space list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSpaceListWithContext(ctx context.Context, request *DescribeSpaceListRequest) (response *DescribeSpaceListResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSpaceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemVariableListRequest() (request *DescribeSystemVariableListRequest) {
    request = &DescribeSystemVariableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeSystemVariableList")
    
    
    return
}

func NewDescribeSystemVariableListResponse() (response *DescribeSystemVariableListResponse) {
    response = &DescribeSystemVariableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemVariableList
// Obtain system variables
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSystemVariableList(request *DescribeSystemVariableListRequest) (response *DescribeSystemVariableListResponse, err error) {
    return c.DescribeSystemVariableListWithContext(context.Background(), request)
}

// DescribeSystemVariableList
// Obtain system variables
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSystemVariableListWithContext(ctx context.Context, request *DescribeSystemVariableListRequest) (response *DescribeSystemVariableListResponse, err error) {
    if request == nil {
        request = NewDescribeSystemVariableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeSystemVariableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemVariableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemVariableListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageSummaryListRequest() (request *DescribeUsageSummaryListRequest) {
    request = &DescribeUsageSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeUsageSummaryList")
    
    
    return
}

func NewDescribeUsageSummaryListResponse() (response *DescribeUsageSummaryListResponse) {
    response = &DescribeUsageSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageSummaryList
// This API is used to query aggregated details of resource usage, including three types of resources: models, plug-ins, and platforms. It displays indicators such as call count and Token consumption, aggregated by space or application dimension.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUsageSummaryList(request *DescribeUsageSummaryListRequest) (response *DescribeUsageSummaryListResponse, err error) {
    return c.DescribeUsageSummaryListWithContext(context.Background(), request)
}

// DescribeUsageSummaryList
// This API is used to query aggregated details of resource usage, including three types of resources: models, plug-ins, and platforms. It displays indicators such as call count and Token consumption, aggregated by space or application dimension.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUsageSummaryListWithContext(ctx context.Context, request *DescribeUsageSummaryListRequest) (response *DescribeUsageSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeUsageSummaryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeUsageSummaryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVariableRequest() (request *DescribeVariableRequest) {
    request = &DescribeVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeVariable")
    
    
    return
}

func NewDescribeVariableResponse() (response *DescribeVariableResponse) {
    response = &DescribeVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVariable
// Get parameter variables
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVariable(request *DescribeVariableRequest) (response *DescribeVariableResponse, err error) {
    return c.DescribeVariableWithContext(context.Background(), request)
}

// DescribeVariable
// Get parameter variables
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVariableWithContext(ctx context.Context, request *DescribeVariableRequest) (response *DescribeVariableResponse, err error) {
    if request == nil {
        request = NewDescribeVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVariableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVariableListRequest() (request *DescribeVariableListRequest) {
    request = &DescribeVariableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeVariableList")
    
    
    return
}

func NewDescribeVariableListResponse() (response *DescribeVariableListResponse) {
    response = &DescribeVariableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVariableList
// Gets the parameter variable list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVariableList(request *DescribeVariableListRequest) (response *DescribeVariableListResponse, err error) {
    return c.DescribeVariableListWithContext(context.Background(), request)
}

// DescribeVariableList
// Gets the parameter variable list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVariableListWithContext(ctx context.Context, request *DescribeVariableListRequest) (response *DescribeVariableListResponse, err error) {
    if request == nil {
        request = NewDescribeVariableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeVariableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVariableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVariableListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowRunRequest() (request *DescribeWorkflowRunRequest) {
    request = &DescribeWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "DescribeWorkflowRun")
    
    
    return
}

func NewDescribeWorkflowRunResponse() (response *DescribeWorkflowRunResponse) {
    response = &DescribeWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowRun
// Queries workflow running details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWorkflowRun(request *DescribeWorkflowRunRequest) (response *DescribeWorkflowRunResponse, err error) {
    return c.DescribeWorkflowRunWithContext(context.Background(), request)
}

// DescribeWorkflowRun
// Queries workflow running details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWorkflowRunWithContext(ctx context.Context, request *DescribeWorkflowRunRequest) (response *DescribeWorkflowRunResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "DescribeWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewExportQARequest() (request *ExportQARequest) {
    request = &ExportQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ExportQA")
    
    
    return
}

func NewExportQAResponse() (response *ExportQAResponse) {
    response = &ExportQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportQA
// Asynchronously export QAs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportQA(request *ExportQARequest) (response *ExportQAResponse, err error) {
    return c.ExportQAWithContext(context.Background(), request)
}

// ExportQA
// Asynchronously export QAs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportQAWithContext(ctx context.Context, request *ExportQARequest) (response *ExportQAResponse, err error) {
    if request == nil {
        request = NewExportQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ExportQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportQAResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentRequest() (request *ModifyAgentRequest) {
    request = &ModifyAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyAgent")
    
    
    return
}

func NewModifyAgentResponse() (response *ModifyAgentResponse) {
    response = &ModifyAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgent
// Modify Agent configuration information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAgent(request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    return c.ModifyAgentWithContext(context.Background(), request)
}

// ModifyAgent
// Modify Agent configuration information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAgentWithContext(ctx context.Context, request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    if request == nil {
        request = NewModifyAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApp
// Modifies applications
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// Modifies applications
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppTriggerRequest() (request *ModifyAppTriggerRequest) {
    request = &ModifyAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyAppTrigger")
    
    
    return
}

func NewModifyAppTriggerResponse() (response *ModifyAppTriggerResponse) {
    response = &ModifyAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAppTrigger
// ModifyAppTrigger
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAppTrigger(request *ModifyAppTriggerRequest) (response *ModifyAppTriggerResponse, err error) {
    return c.ModifyAppTriggerWithContext(context.Background(), request)
}

// ModifyAppTrigger
// ModifyAppTrigger
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAppTriggerWithContext(ctx context.Context, request *ModifyAppTriggerRequest) (response *ModifyAppTriggerResponse, err error) {
    if request == nil {
        request = NewModifyAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConversationRequest() (request *ModifyConversationRequest) {
    request = &ModifyConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyConversation")
    
    
    return
}

func NewModifyConversationResponse() (response *ModifyConversationResponse) {
    response = &ModifyConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConversation
// Modify session information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyConversation(request *ModifyConversationRequest) (response *ModifyConversationResponse, err error) {
    return c.ModifyConversationWithContext(context.Background(), request)
}

// ModifyConversation
// Modify session information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyConversationWithContext(ctx context.Context, request *ModifyConversationRequest) (response *ModifyConversationResponse, err error) {
    if request == nil {
        request = NewModifyConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConversationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocListRequest() (request *ModifyDocListRequest) {
    request = &ModifyDocListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyDocList")
    
    
    return
}

func NewModifyDocListResponse() (response *ModifyDocListResponse) {
    response = &ModifyDocListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocList
// Batch modify documents.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyDocList(request *ModifyDocListRequest) (response *ModifyDocListResponse, err error) {
    return c.ModifyDocListWithContext(context.Background(), request)
}

// ModifyDocList
// Batch modify documents.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyDocListWithContext(ctx context.Context, request *ModifyDocListRequest) (response *ModifyDocListResponse, err error) {
    if request == nil {
        request = NewModifyDocListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyDocList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKBRequest() (request *ModifyKBRequest) {
    request = &ModifyKBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyKB")
    
    
    return
}

func NewModifyKBResponse() (response *ModifyKBResponse) {
    response = &ModifyKBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKB
// This API is used to modify the knowledge base.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyKB(request *ModifyKBRequest) (response *ModifyKBResponse, err error) {
    return c.ModifyKBWithContext(context.Background(), request)
}

// ModifyKB
// This API is used to modify the knowledge base.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyKBWithContext(ctx context.Context, request *ModifyKBRequest) (response *ModifyKBResponse, err error) {
    if request == nil {
        request = NewModifyKBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyKB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKB require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKBResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLabelRequest() (request *ModifyLabelRequest) {
    request = &ModifyLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyLabel")
    
    
    return
}

func NewModifyLabelResponse() (response *ModifyLabelResponse) {
    response = &ModifyLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLabel
// Modify a tag
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLabel(request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
    return c.ModifyLabelWithContext(context.Background(), request)
}

// ModifyLabel
// Modify a tag
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLabelWithContext(ctx context.Context, request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
    if request == nil {
        request = NewModifyLabelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyLabel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// Modify single QA
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// Modify single QA
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyQAWithContext(ctx context.Context, request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    if request == nil {
        request = NewModifyQARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyQA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQAListRequest() (request *ModifyQAListRequest) {
    request = &ModifyQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyQAList")
    
    
    return
}

func NewModifyQAListResponse() (response *ModifyQAListResponse) {
    response = &ModifyQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQAList
// Batch modify QAs
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyQAList(request *ModifyQAListRequest) (response *ModifyQAListResponse, err error) {
    return c.ModifyQAListWithContext(context.Background(), request)
}

// ModifyQAList
// Batch modify QAs
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyQAListWithContext(ctx context.Context, request *ModifyQAListRequest) (response *ModifyQAListResponse, err error) {
    if request == nil {
        request = NewModifyQAListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyQAList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAListResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpaceRequest() (request *ModifySpaceRequest) {
    request = &ModifySpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifySpace")
    
    
    return
}

func NewModifySpaceResponse() (response *ModifySpaceResponse) {
    response = &ModifySpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySpace
// Edit a space
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifySpace(request *ModifySpaceRequest) (response *ModifySpaceResponse, err error) {
    return c.ModifySpaceWithContext(context.Background(), request)
}

// ModifySpace
// Edit a space
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifySpaceWithContext(ctx context.Context, request *ModifySpaceRequest) (response *ModifySpaceResponse, err error) {
    if request == nil {
        request = NewModifySpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifySpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVariableRequest() (request *ModifyVariableRequest) {
    request = &ModifyVariableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ModifyVariable")
    
    
    return
}

func NewModifyVariableResponse() (response *ModifyVariableResponse) {
    response = &ModifyVariableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVariable
// Update parameter variables
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVariable(request *ModifyVariableRequest) (response *ModifyVariableResponse, err error) {
    return c.ModifyVariableWithContext(context.Background(), request)
}

// ModifyVariable
// Update parameter variables
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyVariableWithContext(ctx context.Context, request *ModifyVariableRequest) (response *ModifyVariableResponse, err error) {
    if request == nil {
        request = NewModifyVariableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ModifyVariable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVariable require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVariableResponse()
    err = c.Send(request, response)
    return
}

func NewPauseAppTriggerRequest() (request *PauseAppTriggerRequest) {
    request = &PauseAppTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "PauseAppTrigger")
    
    
    return
}

func NewPauseAppTriggerResponse() (response *PauseAppTriggerResponse) {
    response = &PauseAppTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseAppTrigger
// PauseAppTrigger
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) PauseAppTrigger(request *PauseAppTriggerRequest) (response *PauseAppTriggerResponse, err error) {
    return c.PauseAppTriggerWithContext(context.Background(), request)
}

// PauseAppTrigger
// PauseAppTrigger
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) PauseAppTriggerWithContext(ctx context.Context, request *PauseAppTriggerRequest) (response *PauseAppTriggerResponse, err error) {
    if request == nil {
        request = NewPauseAppTriggerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "PauseAppTrigger")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseAppTrigger require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseAppTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseSkillRequest() (request *ReleaseSkillRequest) {
    request = &ReleaseSkillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ReleaseSkill")
    
    
    return
}

func NewReleaseSkillResponse() (response *ReleaseSkillResponse) {
    response = &ReleaseSkillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseSkill
// This API is used to list a skill.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReleaseSkill(request *ReleaseSkillRequest) (response *ReleaseSkillResponse, err error) {
    return c.ReleaseSkillWithContext(context.Background(), request)
}

// ReleaseSkill
// This API is used to list a skill.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReleaseSkillWithContext(ctx context.Context, request *ReleaseSkillRequest) (response *ReleaseSkillResponse, err error) {
    if request == nil {
        request = NewReleaseSkillRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ReleaseSkill")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseSkill require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseSkillResponse()
    err = c.Send(request, response)
    return
}

func NewResetConversationRequest() (request *ResetConversationRequest) {
    request = &ResetConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "ResetConversation")
    
    
    return
}

func NewResetConversationResponse() (response *ResetConversationResponse) {
    response = &ResetConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetConversation
// Reset a session
//
// Note: Resetting is not supported for the current Claw mode application session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConversation(request *ResetConversationRequest) (response *ResetConversationResponse, err error) {
    return c.ResetConversationWithContext(context.Background(), request)
}

// ResetConversation
// Reset a session
//
// Note: Resetting is not supported for the current Claw mode application session.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConversationWithContext(ctx context.Context, request *ResetConversationRequest) (response *ResetConversationResponse, err error) {
    if request == nil {
        request = NewResetConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "ResetConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetConversationResponse()
    err = c.Send(request, response)
    return
}

func NewRetryReleaseRequest() (request *RetryReleaseRequest) {
    request = &RetryReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "RetryRelease")
    
    
    return
}

func NewRetryReleaseResponse() (response *RetryReleaseResponse) {
    response = &RetryReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRelease
// Retry publishing (republish after suspension)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryRelease(request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    return c.RetryReleaseWithContext(context.Background(), request)
}

// RetryRelease
// Retry publishing (republish after suspension)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryReleaseWithContext(ctx context.Context, request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    if request == nil {
        request = NewRetryReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "RetryRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackReleaseRequest() (request *RollbackReleaseRequest) {
    request = &RollbackReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "RollbackRelease")
    
    
    return
}

func NewRollbackReleaseResponse() (response *RollbackReleaseResponse) {
    response = &RollbackReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackRelease
// Roll back a release
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RollbackRelease(request *RollbackReleaseRequest) (response *RollbackReleaseResponse, err error) {
    return c.RollbackReleaseWithContext(context.Background(), request)
}

// RollbackRelease
// Roll back a release
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RollbackReleaseWithContext(ctx context.Context, request *RollbackReleaseRequest) (response *RollbackReleaseResponse, err error) {
    if request == nil {
        request = NewRollbackReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "RollbackRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewRunAppTriggerNowRequest() (request *RunAppTriggerNowRequest) {
    request = &RunAppTriggerNowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "RunAppTriggerNow")
    
    
    return
}

func NewRunAppTriggerNowResponse() (response *RunAppTriggerNowResponse) {
    response = &RunAppTriggerNowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunAppTriggerNow
// RunAppTriggerNow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunAppTriggerNow(request *RunAppTriggerNowRequest) (response *RunAppTriggerNowResponse, err error) {
    return c.RunAppTriggerNowWithContext(context.Background(), request)
}

// RunAppTriggerNow
// RunAppTriggerNow
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunAppTriggerNowWithContext(ctx context.Context, request *RunAppTriggerNowRequest) (response *RunAppTriggerNowResponse, err error) {
    if request == nil {
        request = NewRunAppTriggerNowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "RunAppTriggerNow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunAppTriggerNow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunAppTriggerNowResponse()
    err = c.Send(request, response)
    return
}

func NewSearchKnowledgeRequest() (request *SearchKnowledgeRequest) {
    request = &SearchKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "SearchKnowledge")
    
    
    return
}

func NewSearchKnowledgeResponse() (response *SearchKnowledgeResponse) {
    response = &SearchKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchKnowledge
// Knowledge retrieval
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SearchKnowledge(request *SearchKnowledgeRequest) (response *SearchKnowledgeResponse, err error) {
    return c.SearchKnowledgeWithContext(context.Background(), request)
}

// SearchKnowledge
// Knowledge retrieval
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SearchKnowledgeWithContext(ctx context.Context, request *SearchKnowledgeRequest) (response *SearchKnowledgeResponse, err error) {
    if request == nil {
        request = NewSearchKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "SearchKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchKnowledgeResponse()
    err = c.Send(request, response)
    return
}
